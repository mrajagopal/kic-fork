HUGO?=hugo
HUGO_IMG?=hugomods/hugo:0.134.1

THEME_MODULE = github.com/nginxinc/nginx-hugo-theme
THEME_VERSION ?= $(shell curl -s https://api.github.com/repos/nginxinc/nginx-hugo-theme/releases/latest | jq -r ".tag_name")

ifeq (, $(shell ${HUGO} version 2> /dev/null))
ifeq (, $(shell docker version 2> /dev/null))
    $(error Docker and Hugo are not installed. Hugo (<0.91) or Docker are required to build the local preview.)
else
    HUGO=docker run --rm -it -v ${CURDIR}:/src -p 1313:1313 ${HUGO_IMG} hugo --bind 0.0.0.0 -p 1313
endif
endif

HUGO_CI=docker run --rm -v ${CURDIR}:/src ${HUGO_IMG} hugo

MARKDOWNLINT?=markdownlint
MARKDOWNLINT_IMG?=ghcr.io/igorshubovych/markdownlint-cli:latest

ifeq (, $(shell ${MARKDOWNLINT} version 2> /dev/null))
ifeq (, $(shell docker version 2> /dev/null))
    $(error Docker and markdownlint are not installed. markdownlint or Docker are required to lint.)
endif
else
    MARKDOWNLINT=docker run --rm -i -v ${CURDIR}:/src --workdir /src ${MARKDOWNLINT_IMG}
endif

MARKDOWNLINKCHECK?=markdown-link-check
MARKDOWNLINKCHECK_IMG?=ghcr.io/tcort/markdown-link-check:stable

ifeq (, $(shell ${MARKDOWNLINKCHECK} --version 2> /dev/null))
ifeq (, $(shell docker version 2> /dev/null))
    $(error Docker and markdown-link-check are not installed. markdown-link-check or Docker are required to check links.)
endif
else
    MARKDOWNLINKCHECK=docker run --rm -it -v ${CURDIR}:/site --workdir /site ${MARKDOWNLINKCHECK_IMG}
endif

.PHONY: docs docs-draft docs-local clean hugo-get hugo-tidy lint-markdown link-check

docs:
	${HUGO}

docs-ci:
	${HUGO_CI}

watch:
	${HUGO} --bind 0.0.0.0 -p 1313 server --disableFastRender

drafts:
	${HUGO} --bind 0.0.0.0 -p 1313 server -D --disableFastRender

clean:
	hugo mod clean
	[ -d "public" ] && rm -rf "public"

hugo-get:
	hugo mod get $(THEME_MODULE)@$(THEME_VERSION)

hugo-tidy:
	hugo mod tidy

hugo-update: hugo-get hugo-tidy

lint-markdown:
	${MARKDOWNLINT} -c .markdownlint.yaml  -- content

link-check:
	${MARKDOWNLINKCHECK} $(shell find content -name '*.md')
