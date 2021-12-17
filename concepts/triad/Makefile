render:
	scripts/triad.py render

ponoko:
	./triad.sh generate-ponoko

perimeters:
	./triad.sh generate-perimeters

kicad_pcb:
	./triad.sh generate-pcb

traces:
	./triad.sh generate-traces

expand_vertices: expand_vertices_left expand_vertices_center expand_vertices_right

expand_vertices_left:
	python3 "./scripts/triad.py" expand-vertices \
		"./temp/pcb_edges_left.json" \
		"./temp/pcb_edges_left_expanded.json" \
		-3.0 \
		--debug-svg="./temp/pcb_edges_left_expanded_debug.svg"

expand_vertices_center:
	python3 "./scripts/triad.py" expand-vertices \
		"./temp/pcb_edges_center.json" \
		"./temp/pcb_edges_center_expanded.json" \
		-3.0 \
		--debug-svg="./temp/pcb_edges_center_expanded_debug.svg"

expand_vertices_right:
	python3 "./scripts/triad.py" expand-vertices \
		"./temp/pcb_edges_right.json" \
		"./temp/pcb_edges_right_expanded.json" \
		3.0 \
		--debug-svg="./temp/pcb_edges_right_expanded_debug.svg"

generate_curves: generate_curves_left generate_curves_center generate_curves_right

generate_curves_left:
	python3 "./scripts/triad.py" generate-curves \
		"./temp/pcb_edges_left_expanded.json" \
		"./temp/pcb_edges_left_curves.json" \
		3.0 \
		--debug-svg="./temp/pcb_edges_left_curves_debug.svg"

generate_curves_center:
	python3 "./scripts/triad.py" generate-curves \
		"./temp/pcb_edges_center_expanded.json" \
		"./temp/pcb_edges_center_curves.json" \
		3.0 \
		--debug-svg="./temp/pcb_edges_center_curves_debug.svg"

generate_curves_right:
	python3 "./scripts/triad.py" generate-curves \
		"./temp/pcb_edges_right_expanded.json" \
		"./temp/pcb_edges_right_curves.json" \
		3.0 \
		--debug-svg="./temp/pcb_edges_right_curves_debug.svg"
