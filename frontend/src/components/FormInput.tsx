import { HTMLInputTypeAttribute } from "react";
import { Control, Controller } from "react-hook-form";

interface FormInputProps {
  type: HTMLInputTypeAttribute;
  label?: string;
  message?: string;
  name: string;
  control: Control<any>;
  required: boolean;
}

export default function FormInput(props: FormInputProps) {
  return (
    <div className="form-control w-full max-w-xs">
      {props.label && (
        <label className="label">
          <span className="text-bold">{props.label}</span>
        </label>
      )}
      <Controller
        name={props.name}
        control={props.control}
        render={({ field: { onChange, onBlur, value } }) => (
          <input
            type={props.type}
            placeholder="Type here"
            className="input input-bordered w-full max-w-xs"
            onChange={onChange}
            onBlur={onBlur}
            value={value}
            required={props.required}
          />
        )}
      />
      {props.message && (
        <label className="label">
          <span className="label-text">{props.message}</span>
        </label>
      )}
    </div>
  );
}
