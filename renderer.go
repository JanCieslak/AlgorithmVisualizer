package main

import (
	"github.com/go-gl/gl/all-core/gl"
)

type Renderer struct {
	vao uint32
	vbo uint32
	vertices []float32
}

const MaxVertices = 100000

func newRenderer() *Renderer {
	renderer := &Renderer {}

	gl.GenVertexArrays(1, &renderer.vao)
	gl.GenBuffers(1, &renderer.vbo)

	gl.BindVertexArray(renderer.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, renderer.vbo)

	gl.BufferData(gl.ARRAY_BUFFER, MaxVertices, nil, gl.DYNAMIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 5 * 4, nil)

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 2, gl.FLOAT, false, 5 * 4, gl.PtrOffset(8))

	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 1, gl.FLOAT, false, 5 * 4, gl.PtrOffset(16))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return renderer
}

func (it *Renderer) Begin()  {
	it.vertices = it.vertices[:0]
}

func (it *Renderer) Add(vertices []float32)  {
	it.vertices = append(it.vertices, vertices...)
}

func (it *Renderer) End()  {
	gl.BindVertexArray(it.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, it.vbo)

	gl.BufferSubData(gl.ARRAY_BUFFER, 0, 4 * len(it.vertices), gl.Ptr(it.vertices))
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(it.vertices)))

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (it *Renderer) AddGrid(grid *Grid) {
	for row := 0; row < grid.rows; row++ {
		for col := 0; col < grid.cols; col++ {
			it.Add([]float32 {
				float32(col 	  * grid.cellWidth), float32(row 	   * grid.cellHeight), 0, 0, float32(grid.cells[row][col]),
				float32((col + 1) * grid.cellWidth), float32((row + 1) * grid.cellHeight), 1, 1, float32(grid.cells[row][col]),
				float32(col 	  * grid.cellWidth), float32((row + 1) * grid.cellHeight), 0, 1, float32(grid.cells[row][col]),
				float32(col  	  * grid.cellWidth), float32(row 	   * grid.cellHeight), 0, 0, float32(grid.cells[row][col]),
				float32((col + 1) * grid.cellWidth), float32((row + 1) * grid.cellHeight), 1, 1, float32(grid.cells[row][col]),
				float32((col + 1) * grid.cellWidth), float32(row  	   * grid.cellHeight), 1, 0, float32(grid.cells[row][col]),
			})
		}
	}
}