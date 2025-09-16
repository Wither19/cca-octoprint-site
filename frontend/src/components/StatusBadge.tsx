type StatusBadgeProps = { status: string, color: string };

function StatusBadge({ status, color }: StatusBadgeProps) {
  return (
     <span className={`badge bg-${color} fs-6 position-relative`}>{status}</span>
  )
}

export default StatusBadge;