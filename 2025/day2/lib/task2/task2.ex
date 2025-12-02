defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> Enum.flat_map(&get_invalid_ids/1)
    |> Enum.sum()
  end

  defp parse(input) do
    String.trim(input)
    |> String.split(",")
    |> Enum.map(fn x ->
      String.split(x, "-")
      |> Enum.map(fn x ->
        {i, _} = Integer.parse(x)
        i
      end)
      |> List.to_tuple()
    end)
  end

  defp get_invalid_ids({l, r}) do
    get_invalid_ids({l, r}, [])
  end

  defp get_invalid_ids({l, r}, acc) when l > r do
    acc
  end

  defp get_invalid_ids({l, r}, acc) do
    new_acc =
      case invalid_id?(l) do
        true ->
          [l] ++ acc

        false ->
          acc
      end

    next_l = get_next_id(l)
    get_invalid_ids({next_l, r}, new_acc)
  end

  def digits_total(num) do
    floor(:math.log10(num)) + 1
  end

  # the main change from Task1:
  def invalid_id?(num) do
    num_digits = digits_total(num)

    case num_digits < 2 do
      true ->
        false

      false ->
        chars = Integer.to_charlist(num)

        Range.new(1, div(num_digits, 2))
        |> Enum.any?(fn x ->
          Enum.chunk_every(chars, x)
          |> Enum.uniq()
          |> Enum.count() <= 1
        end)
    end
  end

  # mostly bruteforce, but optimized to skip ranges of odd total number of digits to neareset possible invalid id
  defp get_next_id(num) do
    num + 1
  end
end
