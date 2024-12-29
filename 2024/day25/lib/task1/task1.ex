defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {locks, keys} =
      input
      |> parse()

    Enum.map(locks, fn lock ->
      Enum.filter(keys, fn key -> can_open?(lock, key) end)
      |> length()
    end)
    |> Enum.sum()
  end

  defp can_open?(lock, key) do
    llock = Tuple.to_list(lock)
    lkey = Tuple.to_list(key)

    Enum.zip(llock, lkey)
    |> Enum.reduce(true, fn {l, r}, acc ->
      l + r <= 5 and acc
    end)
  end

  defp parse(input) do
    inputs =
      input
      |> String.split("\n", trim: true)
      # base of lock/key + 6 rows (0 to 5)
      |> Enum.chunk_every(7)

    locks =
      inputs
      |> Enum.filter(fn
        ["#####" | _k] -> true
        _ -> false
      end)
      |> Enum.map(fn [_header | lock] -> parse_entity(lock) end)

    keys =
      inputs
      |> Enum.map(fn x -> Enum.reverse(x) end)
      |> Enum.filter(fn
        ["#####" | _k] -> true
        _ -> false
      end)
      |> Enum.map(fn [_header | key] -> parse_entity(key) end)

    {locks, keys}
  end

  defp parse_entity(e) do
    Enum.reduce(e, {0, 0, 0, 0, 0}, fn x, acc ->
      x
      |> String.to_charlist()
      |> add_to_tuple(acc)
    end)
  end

  defp add_to_tuple(list, tuple) when is_list(list) do
    list
    |> Enum.with_index()
    |> Enum.reduce(tuple, fn {el, idx}, acc ->
      case el do
        ?# ->
          put_elem(acc, idx, elem(acc, idx) + 1)

        _ ->
          acc
      end
    end)
  end
end
