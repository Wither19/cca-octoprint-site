import type { PrinterState } from "../types"

import { printerStateColors } from "../functions"


type StatusBadgeProps = { state: PrinterState }

function StatusBadge({ state }: StatusBadgeProps) {
  return (
     <span className={`badge bg-${printerStateColors(state.Flags)} fs-6 position-relative`}>{state.Text}</span>
  )
}

export default StatusBadge