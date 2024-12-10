defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {_, _, spaces, files} =
      input
      |> String.trim()
      |> String.to_charlist()
      |> Enum.map(&(&1 - ?0))
      |> Enum.with_index()
      |> Enum.reduce({0, 0, [], []}, fn {x, i}, {pos, file_id, spaces, files} = acc ->
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
              {pos + x, file_id + 1, spaces,
               [{pos, x, file_id}, files]}

            :space ->
              {pos + x, file_id, [{pos, x}, spaces], files}
          end
        end
      end)

    # move(spaces |> Enum.reverse(), files)
  end

  # defp move(_, []) do
  #   moved_files
  # end

  # defp move(spaces, [{from,to,fid}]) do
  #   {from, to, fid} = Enum.at(0)
  #   new_files = Enum.drop(files, 1)

  #   # store spaces in min heap by pos
  #   # iterate over spaces to find first where val > needed space
  #   # remove from min_heap
  #   # put file into moved_files
  #   # if any space left over put back into min_heap
  #   # if file not moved - copy original file into moved_files


  #   move(new_spaces, new_files)
  # end
end
