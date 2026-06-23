package core

import "strings"

var (
	ColorGreen  = "\033[38;2;103;156;116m"
	ColorRed    = "\033[38;2;222;122;122m"
	ColorBlue   = "\033[38;2;98;149;217m"
	ColorYellow = "\033[38;2;241;229;121m"
	ColorPink   = "\033[38;2;234;128;252m"
	ColorReset  = "\033[0m"

	ColorGreenHex  = "[#679c74]"
	ColorRedHex    = "[#de7a7a]"
	ColorBlueHex   = "[#6295d9]"
	ColorYellowHex = "[#f1e579]"
	ColorPinkHex   = "[#ea80fc]"
	ColorResetHex  = "[-]"
)

func GetColor(msg string, hex bool) string {
	lines := strings.SplitN(msg, "\n", 2)
	if len(lines) == 0 {
		return msg
	}

	line := strings.TrimSpace(lines[0])
	parts := strings.SplitN(line, "|", 2)
	if len(parts) >= 2 && strings.Contains(parts[0], "->") {
		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])

		fromTo := strings.SplitN(left, "->", 2)
		if len(fromTo) >= 2 {
			from := strings.TrimSpace(fromTo[0])
			to := strings.TrimSpace(fromTo[1])

			yellow := ColorYellow
			pink := ColorPink
			blue := ColorBlue
			green := ColorGreen
			reset := ColorReset

			if hex {
				yellow = ColorYellowHex
				pink = ColorPinkHex
				blue = ColorBlueHex
				green = ColorGreenHex
				reset = ColorResetHex
			}

			coloredLine := yellow + from + reset
			coloredLine += " -> " + pink + to + reset
			coloredLine += " | " + blue + right + reset

			if len(lines) > 1 {
				text := strings.TrimPrefix(lines[1], ": ")
				coloredLine += "\n: " + green + text + reset
			}

			return coloredLine
		}
	}

	if strings.HasPrefix(line, "Error:") || strings.HasPrefix(line, "Unknown command") {
		if hex {
			return ColorRedHex + msg + ColorResetHex
		}
		return ColorRed + msg + ColorReset
	}

	return msg
}
