defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {rules, updates} =
      input
      |> String.split("\n", trim: true)
      |> Enum.split_with(fn x -> String.contains?(x, "|") end)

    all_rules =
      rules
      |> parse_rules()

    updates
    |> Enum.map(fn x -> String.split(x, ",") |> Enum.map(&String.to_integer/1) end)
    |> Enum.map(fn x ->
      ordered = order_for(all_rules, x)

      if not ordered?(ordered, x) do
        Enum.at(ordered, div(length(ordered), 2))
      else
        0
      end
    end)
    |> Enum.sum()
  end

  defp remove_contains({l, r}, el) do
    to_delete_l =
      Enum.with_index(l)
      |> Enum.filter(fn
        {^el, _} ->
          true

        {_, _} ->
          false
      end)
      |> Enum.map(fn {_, i} -> i end)

    to_delete_r =
      Enum.with_index(r)
      |> Enum.filter(fn
        {^el, _} ->
          true

        {_, _} ->
          false
      end)
      |> Enum.map(fn {_, i} -> i end)

    deletable = Enum.uniq(to_delete_l ++ to_delete_r)

    new_l =
      Enum.with_index(l)
      |> Enum.reject(fn {_, i} -> i in deletable end)
      |> Enum.map(&elem(&1, 0))

    new_r =
      Enum.with_index(r)
      |> Enum.reject(fn {_, i} -> i in deletable end)
      |> Enum.map(&elem(&1, 0))

    {new_l, new_r}
  end

  defp remove_not_contains({l, r}, elements) when is_list(elements) do
    to_delete_l =
      Enum.with_index(l)
      |> Enum.filter(fn
        {el, _} ->
          el not in elements
      end)
      |> Enum.map(fn {_, i} -> i end)

    to_delete_r =
      Enum.with_index(r)
      |> Enum.filter(fn
        {el, _} ->
          el not in elements
      end)
      |> Enum.map(fn {_, i} -> i end)

    deletable = Enum.uniq(to_delete_l ++ to_delete_r)

    new_l =
      Enum.with_index(l)
      |> Enum.reject(fn {_, i} -> i in deletable end)
      |> Enum.map(&elem(&1, 0))

    new_r =
      Enum.with_index(r)
      |> Enum.reject(fn {_, i} -> i in deletable end)
      |> Enum.map(&elem(&1, 0))

    {new_l, new_r}
  end

  defp order_for({l, r}, elements) do
    {new_l, new_r} = remove_not_contains({l, r}, elements)
    order_rules({new_l, new_r})
  end

  defp ordered?([], []) do
    true
  end

  defp ordered?(_, []) do
    true
  end

  defp ordered?([], _) do
    false
  end

  defp ordered?([ho | order], [ht | update]) do
    cond do
      ho == ht ->
        ordered?(order, update)

      ho != ht ->
        ordered?(order, [ht | update])
    end
  end

  defp parse_rules(rules) do
    rules
    |> Enum.map(fn x ->
      [l, r] = String.split(x, "|")
      {String.to_integer(l), String.to_integer(r)}
    end)
    |> Enum.reduce(
      {[], []},
      fn {l, r}, {lacc, racc} -> {[l | lacc], [r | racc]} end
    )
  end

  defp order_rules({[], []}) do
    []
  end

  defp order_rules({l, r}) do
    lx = Enum.uniq(l) |> MapSet.new()
    rx = Enum.uniq(r) |> MapSet.new()
    [first] = MapSet.difference(lx, rx) |> MapSet.to_list()

    if length(rx |> MapSet.to_list()) == 1 do
      [first | MapSet.to_list(rx)]
    else
      {new_l, new_r} = remove_contains({l, r}, first)
      [first | order_rules({new_l, new_r})]
    end
  end
end
