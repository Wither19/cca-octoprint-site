import type { PrinterStatus } from "@jamesgopsill/octoprint-client";

type PrinterStateFlags = PrinterStatus["state"]["flags"]

type Optional<T> = T | null | undefined

export function printerStateColors(s: Optional<PrinterStateFlags>) {
	let color = "secondary";
	if (s) {
		if (s.ready && s.sdReady) {
			color = "success";
		} else if (s.paused || s.printing || s.cancelling) {
			color = "warning";
		} else if (s.error || s.closedOnError) {
			color = "danger";
		}
	}

	return color;
}