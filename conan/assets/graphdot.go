package assets

// TODO: to go-template
var DotContent = `\
digraph {
    {%- for src, dst in graph.edges %}
        "{{ src.label }}" -> "{{ dst.label }}"
    {%- endfor %}
}
`
