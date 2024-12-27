package printer

import (
	"strings"

	. "github.com/aspizu/compiler/external"
	"github.com/aspizu/compilerternal/sourcemap"
	"github.com/aspizu/compilerternal/transform"
)

type PrintCSSResult struct {
	Output         [][]byte
	SourceMapChunk sourcemap.Chunk
}

func PrintCSS(sourcetext string, doc *Node, opts transform.TransformOptions) PrintCSSResult {
	p := &printer{
		opts:    opts,
		builder: sourcemap.MakeChunkBuilder(nil, sourcemap.GenerateLineOffsetTables(sourcetext, len(strings.Split(sourcetext, "\n")))),
	}

	result := PrintCSSResult{
		SourceMapChunk: p.builder.GenerateChunk(p.output),
	}

	if len(doc.Styles) > 0 {
		for _, style := range doc.Styles {
			if style.FirstChild != nil && strings.TrimSpace(style.FirstChild.Data) != "" {
				p.addSourceMapping(style.Loc[0])
				p.print(strings.TrimSpace(style.FirstChild.Data))
				result.Output = append(result.Output, p.output)
				p.output = []byte{}
				p.addNilSourceMapping()
			}
		}
	}

	return result
}
