import { IconAlertTriangle } from "@tabler/icons-react";

interface AlertProps {
  message: string;
  className?: string;
}

export default function Alert(props: AlertProps) {
  return props.message ? (
    <div className={`alert ${props.className}`}>
      <div>
        <IconAlertTriangle />
        <span>{props.message}</span>
      </div>
    </div>
  ) : null;
}
