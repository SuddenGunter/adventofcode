defmodule Task1 do
  @spec solution([String.t()]) :: integer()
  def solution(contents) do
    contents
    # drop cd
    |> Enum.drop(1)
    |> build_tree()
    |> count_dirs_size()
    |> dirs_below()
    |> sum_size()
  end

  @spec build_tree([String.t()], String.t()) :: %Fs.Dir{}
  defp build_tree(lines, dir_name \\ "/") do
    # drop ls
    to_process = Enum.drop(lines, 1)
    cur_dir = %Fs.Dir{name: dir_name}
    children = parse_children(to_process)
    %{cur_dir | children: children}
  end

  defp parse_children(lines) do
    Enum.reduce_while(lines, [], fn x, acc ->
      if String.starts_with?(x, "$ cd") do
        {:halt, acc}
      else
        case String.split(x) do
          ["dir", dir_name] ->
            {:cont,
             [
               build_tree(
                 find_dir(lines, dir_name) |> Enum.drop(1),
                 dir_name
               )
               | acc
             ]}

          ["$", _] ->
            {:halt, acc}

          [f_size, f_name] ->
            {fsize_parsed, _} = Integer.parse(f_size)
            {:cont, [%Fs.File{name: f_name, size: fsize_parsed} | acc]}
        end
      end
    end)
  end

  defp find_dir(lines, dir_name) do
    {drop, _} =
      Enum.reduce_while(lines, {0, 0}, fn l, {drop, depth} ->
        if l == "$ cd #{dir_name}" && depth == 0 do
          {:halt, {drop, depth}}
        else
          cond do
            l == "$ cd .." -> {:cont, {drop + 1, depth - 1}}
            String.starts_with?(l, "$ cd") -> {:cont, {drop + 1, depth + 1}}
            true -> {:cont, {drop + 1, depth}}
          end
        end
      end)

    Enum.drop(lines, drop)
  end

  @spec count_dirs_size(%Fs.Dir{}) :: %Fs.Dir{}
  defp count_dirs_size(root) do
    sized_children =
      Enum.map(root.children, fn x ->
        case x do
          dir when is_struct(x, Fs.Dir) -> count_dirs_size(dir)
          file when is_struct(x, Fs.File) -> file
        end
      end)

    size = sized_children |> Enum.map(fn x -> x.size end) |> Enum.sum()

    Map.merge(root, %Fs.Dir{name: root.name, size: size, children: sized_children})
  end

  @spec dirs_below(%Fs.Dir{}, Integer.t()) :: [Fs.Dir]
  defp dirs_below(root, limit \\ 100_000) do
    res =
      Enum.reduce(root.children, [], fn x, acc ->
        node =
          case x do
            dir when is_struct(x, Fs.Dir) ->
              dirs_below(dir, limit)

            _ when is_struct(x, Fs.File) ->
              []
          end

        [node | acc]
      end)
      |> List.flatten()

    if root.size <= limit do
      [root | res]
    else
      res
    end
  end

  @spec sum_size([Fs.Dir]) :: integer()
  defp sum_size(dirs) do
    Enum.map(dirs, fn x -> x.name end) |> IO.inspect()

    dirs
    |> Enum.map(fn x -> x.size end)
    |> Enum.sum()
  end
end
