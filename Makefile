SOLUTION=src/KbUtil.sln
RUNTIME_WIN="win7-x64"
RUNTIME_LNX="linux-x64"
PUBLISH_OPTS=--self-contained --configuration Release

all: build

publish: publish_windows publish_linux

publish_windows: build
	dotnet publish --runtime $(RUNTIME_WIN) $(PUBLISH_OPTS) $(SOLUTION)

publish_linux: build
	dotnet publish --runtime $(RUNTIME_LNX) $(PUBLISH_OPTS) $(SOLUTION)

build:
	dotnet build -c Debug $(SOLUTION)
	dotnet build -c Release $(SOLUTION)

clean:
	rm -rf build
