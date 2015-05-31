package nbt

func indent(s string) (out string) {
	last := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			out += s[last:i+1] + "	"
			last = i + 1
		}
	}
	out += s[last:]
	return out
}
