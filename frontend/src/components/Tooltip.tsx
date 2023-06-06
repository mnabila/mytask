import { ReactNode } from "react";

interface TooltipProps {
  data: string;
  className?: string;
  children: ReactNode;
}
export default function Tooltip({ className, data, children }: TooltipProps) {
  return (
    <div className={`tooltip ${className}`} data-tip={data}>
      {children}
    </div>
  );
}
