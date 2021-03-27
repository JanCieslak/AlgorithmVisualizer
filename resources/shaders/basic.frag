#version 400 core

in vec2 p_uv;
in float p_textureid;

out vec4 color;

uniform sampler2D tex[16];

void main() {
    color = texture(tex[int(p_textureid)], p_uv);
}