SOLUTION=src/KbUtil.sln
RUNTIME_WIN="win7-x64"
RUNTIME_LNX="linux-x64"
PUBLISH_OPTS=--self-contained --configuration Release

PROJ_KBUTIL=src/KbUtil.Console/KbUtil.Console.csproj

SRC_FILES=$(shell find src/ -type f)

.PHONY: all
all: build

.PHONY: publish
publish:
	dotnet publish --runtime $(RUNTIME_WIN) $(PUBLISH_OPTS) --output "../../build/publish/win7-x64/kbutil" $(PROJ_KBUTIL)
	dotnet publish --runtime $(RUNTIME_LNX) $(PUBLISH_OPTS) --output "../../build/publish/linux-x64/kbutil" $(PROJ_KBUTIL)

.PHONY: build
build: $(SRC_FILES)
	dotnet build -c Debug $(SOLUTION)
	dotnet build -c Release $(SOLUTION)

.PHONY: test
test:
	dotnet test src/KbUtil.Lib.Tests

.PHONY: clean
clean:
	rm -rf build
