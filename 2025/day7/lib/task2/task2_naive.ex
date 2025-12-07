defmodule Task2Naive do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> first_beam()
    |> project_beam()
  end

  defp parse(input) do
    String.trim(input)
    |> String.split("\n")
    |> Enum.with_index()
    |> Enum.map(fn {row, row_number} ->
      String.graphemes(row)
      |> Enum.with_index()
      |> Enum.filter(fn
        {".", _} -> false
        _ -> true
      end)
      |> Enum.map(fn
        {"S", start_pos} -> {:start, {row_number, start_pos}}
        {"^", splitter_pos} -> {:splitter, {row_number, splitter_pos}}
      end)
    end)
  end

  defp first_beam([[{:start, {x, y}}] | tail]) do
    {tail, [{x + 1, y}], 0}
  end

  defp project_beam({[], beams, _total_splits}) do
    length(beams)
  end

  defp project_beam({[[] | tail_cmds], beams, total_splits}) do
    project_beam({tail_cmds, project_beams_to_next_row(beams), total_splits})
  end

  defp project_beam({[row_cmds | tail_cmds], beams, total_splits}) do
    IO.inspect(%{
      row_cmds: row_cmds,
      le: length(beams)
    })

    {projected_beams, new_splits} =
      Enum.reduce(row_cmds, {beams, total_splits}, fn {:splitter, {x, y} = splitter_pos},
                                                      {acc_beams, acc_splits} ->
        hit_splitter = Enum.filter(acc_beams, fn beam_pos -> beam_pos == splitter_pos end)

        cond do
          length(hit_splitter) > 0 ->
            new_beams =
              List.duplicate([{x, y - 1}, {x, y + 1}], length(hit_splitter))
              |> Enum.flat_map(fn x -> x end)

            {(new_beams ++ acc_beams)
             |> Enum.filter(fn
               beam_pos when beam_pos == splitter_pos ->
                 false

               _ ->
                 true
             end), acc_splits + 1}

          length(hit_splitter) == 0 ->
            {acc_beams, acc_splits}
        end
      end)

    project_beam(
      {tail_cmds,
       projected_beams
       |> project_beams_to_next_row(), new_splits}
    )
  end

  defp project_beams_to_next_row(beams) do
    Enum.map(beams, fn {x, y} -> {x + 1, y} end)
  end
end
