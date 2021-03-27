#version 400 core

layout (location = 0) in vec2 a_position;
layout (location = 1) in vec2 a_uv;
layout (location = 2) in float a_textureid;

out vec2 p_uv;
out float p_textureid;

uniform mat4 u_projection;

void main() {
    gl_Position = u_projection * vec4(a_position, 0.0, 1.0);
    p_uv = a_uv;
    p_textureid = a_textureid;
}