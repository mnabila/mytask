import { MouseEventHandler, ReactNode } from "react";

interface ButtonProps {
  icon?: ReactNode;
  label?: string;
  className?: string;
  type?: "button" | "submit" | "reset";
  onClick?: MouseEventHandler<HTMLButtonElement>;
}

export default function Button(props: ButtonProps) {
  return (
    <button
      className={`inline-flex btn ${props.className}`}
      type={props.type}
      onClick={props.onClick}
    >
      {props.icon}
      <span>{props.label}</span>
    </button>
  );
}
