all: cmds 

cmds: pkgs
	cd ./cmd/kbutil && make

pkgs:
	cd ./pkg/kbutil && make

test:
	echo "TODO"

render: pkgs cmds
	mkdir -p ./temp
	build/kbutil testdata/1-2x2-macropad.json 1>./temp/temp.svg

clean:
	rm -rf ./build
