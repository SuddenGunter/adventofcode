defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {_, _, spaces, files} =
      input
      |> String.trim()
      |> String.to_charlist()
      |> Enum.map(&(&1 - ?0))
      |> Enum.with_index()
      |> Enum.reduce({0, 0, Heap.new(), []}, fn {x, i}, {pos, file_id, spaces, files} = acc ->
        type =
          case rem(i, 2) do
            0 -> :file
            1 -> :space
          end

        if x == 0 do
          acc
        else
          case type do
            :file ->
              {pos + x, file_id + 1, spaces, [{pos, x, file_id}] ++ files}

            :space ->
              {pos + x, file_id, Heap.push(spaces, {pos, x}), files}
          end
        end
      end)

    fs = move(spaces, files, [])

    fs |> checksum()
  end

  def checksum(fs) do
    Enum.flat_map(fs, fn {pos, len, val} ->
      for x <- pos..(pos + len - 1), do: val * x
    end)
    |> Enum.sum()
  end

  defp move(_, [], new_fs) do
    new_fs
  end

  defp move(spaces, [{fpos, flen, fid} | tail_files], new_fs) do
    case find_space(spaces, {fpos, flen}, Heap.new()) do
      :not_found ->
        move(spaces, tail_files, [{fpos, flen, fid}] ++ new_fs)

      {spos, new_spaces} ->
        move(new_spaces, tail_files, [{spos, flen, fid}] ++ new_fs)
    end

    # store spaces in min heap by pos
    # iterate over spaces to find first where val > needed space
    # remove from min_heap
    # put file into moved_files
    # if any space left over put back into min_heap
    # if file not moved - copy original file into moved_files
  end

  def find_space(spaces, {pos, len}, new_spaces) do
    case Heap.root(spaces) do
      nil ->
        :not_found

      {k, v} ->
        cond do
          k < pos and v - len == 0 ->
            {k,
             Enum.reduce(Heap.pop(spaces), new_spaces, fn {k, v}, acc ->
               Heap.push(acc, {k, v})
             end)}

          k < pos and v - len > 0 ->
            new_size = v - len
            new_pos = k + len

            {k,
             Enum.reduce(Heap.pop(spaces), new_spaces, fn {k, v}, acc ->
               Heap.push(acc, {k, v})
             end)
             |> Heap.push({new_pos, new_size})}

          true ->
            find_space(Heap.pop(spaces), {pos, len}, Heap.push(new_spaces, {k, v}))
        end
    end
  end
end
