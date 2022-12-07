defmodule Task1 do
  @spec solution([String.t()]) :: integer()
  def solution(contents) do
    contents
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
                 Enum.drop_while(lines, fn l -> l != "$ cd #{dir_name}" end) |> Enum.drop(1),
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
    Enum.reduce(root.children, [], fn x, acc ->
      node =
        case x do
          dir when is_struct(x, Fs.Dir) ->
            if dir.size <= 100_000,
              do: [dir | dirs_below(dir, limit)],
              else: dirs_below(dir, limit)

          _ when is_struct(x, Fs.File) ->
            []
        end

      [node | acc]
    end)
    |> List.flatten()
  end

  @spec sum_size([Fs.Dir]) :: integer()
  defp sum_size(dirs) do
    dirs
    |> Enum.map(fn x -> x.size end)
    |> Enum.sum()
  end
end
