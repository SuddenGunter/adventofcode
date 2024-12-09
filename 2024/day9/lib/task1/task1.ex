defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    input
    |> String.to_charlist()
    |> Enum.map(&(&1 - ?0))
    |> Enum.with_index()
    |> Enum.flat_map(fn {x, i} ->
      type = case rem(i, 2) do
        0 -> :file
        1 -> :space
      end

      generate_fs_entry(type, x)
    end)
    |> Enum.reduce({:file, 0, []}, fn x, {state, num, arr} ->
      case {state, x} do
        {:file, :zero_len_space} -> {:space, num, arr}
        {:file, :file} -> {:file, num, [num] ++ arr}
        {:space, :file} -> {:file, num + 1, [num + 1] ++ arr}
        {_, :space} -> {:space, num, [:space]++arr}
      end
    end)
    |> elem(2)
    |> Enum.reverse()
    |> Enum.map(fn x when is_integer(x) -> Integer.to_string(x)
    :space -> "."
  end)
  end


  defp generate_fs_entry(:file, 0) do
    throw("unexpected")
  end

  defp generate_fs_entry(:space, 0) do
    [:zero_len_space]
  end


  defp generate_fs_entry(type, x) do
    List.duplicate(type, x)
  end
end
