import { IconTrash } from "@tabler/icons-react";
import { Button, Tooltip } from "../../components";
import { TodoResponse, TodoRquest } from "../../dto";
import { useEffect, useState } from "react";
import useAxios from "../../hooks";
import { toast } from "react-hot-toast";
import { SubmitHandler, useForm } from "react-hook-form";

export default function Todo() {
  const axios = useAxios();
  const [todo, setTodo] = useState<TodoResponse[]>([]);

  const getTodo = () => {
    axios
      .get("/todo")
      .then((res) => setTodo(res.data.data))
      .catch((err) => toast.error(err.response.data.error));
  };
  useEffect(getTodo, []);

  const { register, handleSubmit, reset } = useForm<TodoRquest>();

  const onSubmit: SubmitHandler<TodoRquest> = (data) => {
    axios
      .post("/todo", data)
      .then((res) => {
        toast.success(res.data.message);
        reset();
      })
      .catch((err) => toast.error(err.response.data.error))
      .finally(getTodo);
  };

  const onDelete = (id: number) => {
    axios
      .delete(`/todo/${id}`)
      .then((res) => {
        toast.success(res.data.message);
      })
      .catch((err) => toast.error(err.response.data.error))
      .finally(getTodo);
  };

  return (
    <div className="card w-1/3 bg-base-100 shadow-xl">
      <div className="card-body">
        <div className="flex flex-col space-y-3 overflow-y-auto h-96 max-h-96">
          {todo.map((t) => {
            return (
              <div className="flex border rounded-md p-2 bg-base-200">
                <div className="flex-1 m-auto">{t.task}</div>
                <Tooltip data="Hapus" className="tooltip-left">
                  <Button icon={<IconTrash />} onClick={() => onDelete(t.id)} />
                </Tooltip>
              </div>
            );
          })}
        </div>
        <div className="card-actions justify-end mt-5">
          <form onSubmit={handleSubmit(onSubmit)} className="w-full">
            <input
              type="text"
              placeholder="Type here"
              className="input input-bordered w-full"
              {...register("task", { required: true })}
            />
          </form>
        </div>
      </div>
    </div>
  );
}
