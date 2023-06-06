import { ReactNode } from "react";
import { Link, LinkProps } from "react-router-dom";

interface ButtonLinkProps extends LinkProps {
	icon?: ReactNode;
	label?: string;
}

export default function ButtonLink(props: ButtonLinkProps) {
	return (
		<Link className={`inline-flex btn ${props.className}`} to={props.to}>
			{props.icon}
			<span>{props.label}</span>
		</Link>
	);
}
