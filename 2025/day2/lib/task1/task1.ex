defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> Enum.flat_map(&get_invalid_ids/1)
    |> IO.inspect()
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

    # in bruteforce just +1
    next_l = get_next_id(l)
    get_invalid_ids({next_l, r}, new_acc)
  end

  defp invalid_id?(num) do
    num_digits = digits_total(num)

    case rem(num_digits, 2) do
      1 ->
        # number with uneven count of digits cannot be invalid id
        false

      0 ->
        # compare left and right parts of number
        rem(num, floor(:math.pow(10, num_digits / 2))) ===
          div(num, floor(:math.pow(10, num_digits / 2)))
    end
  end

  defp get_next_id(num) do
    num_digits = digits_total(num)

    case rem(num_digits, 2) do
      1 ->
        # number with uneven count of digits cannot be invalid id, so we skip to next decimal place, e.g. 555 -> 1000
        # considering 1000... will never be an invalid id, we go even further, and start with all 1s - smallest "invalid id" possible, e.g. 1111 for 1000
        Range.new(0, num_digits) |> Enum.reduce(0, fn x, acc -> acc + :math.pow(10, x) end)

      0 ->
        num + 1
    end
  end

  defp digits_total(num) do
    floor(:math.log10(num)) + 1
  end
end
