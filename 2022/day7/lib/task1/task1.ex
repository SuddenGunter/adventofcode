defmodule Task1 do
  @spec solution([String.t()]) :: integer()
  def solution(contents) do
    contents
    # skip cd /
    |> Enum.drop(1)
    |> build_tree()
    |> count_dirs_size()
    |> dirs_below(100_000)
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
            [
              build_tree(
                Enum.drop_while(lines, fn l -> l != "$ cd #{dir_name}" end) |> Enum.drop(1),
                dir_name
              )
              | acc
            ]

          ["$", _] ->
            acc

          [f_size, f_name] ->
            {fsize_parsed, _} = Integer.parse(f_size)
            [%Fs.File{name: f_name, size: fsize_parsed} | acc]
        end
      end
    end)
  end

  @spec solution([Fs.Dir | Fs.File]) :: [Fs.Dir]
  defp count_dirs_size(x) do
    x |> Enum.filter(&only_dirs/1) |> Enum.map(fn x -> %{x | size: 10} end)
  end

  defp only_dirs(f) when is_struct(f, Fs.Dir) do
    true
  end

  defp only_dirs(_) do
    false
  end

  @spec solution([Fs.Dir]) :: [Fs.Dir]
  defp dirs_below(dirs, 100_000) do
    dirs
  end

  @spec solution([Fs.Dir]) :: integer()
  defp sum_size(dirs) do
    dirs
    |> Enum.map(fn x -> x.size end)
    |> Enum.sum()
  end
end
