defmodule Task1 do
  alias Collections.DisjointSet
  @spec solution(String.t()) :: integer()
  def solution(input) do
    parse(input)
    |> distances_heap()
    # 10 for demo example in https://adventofcode.com/2025/day/8, 1000 for real data.txt
    |> connect(DisjointSet.new(), 1000)
    |> take_largest_circuits_size(3)
    |> Enum.product()
  end

  defp take_largest_circuits_size(set, num) do
    {parents, _} =
      Map.get(set, :root)
      |> Map.keys()
      |> Enum.uniq()
      |> Enum.reduce({%{}, set}, fn x, {parents, acc_set} ->
        {parent, optimized_set} = DisjointSet.find(acc_set, x)

        {
          Map.update(parents, parent, 1, fn kids -> kids + 1 end),
          optimized_set
        }
      end)

    Map.to_list(parents)
    |> Enum.sort(fn {_, size_l}, {_, size_r} -> size_l > size_r end)
    |> Enum.take(num)
    |> Enum.map(fn {_, x} -> x end)
  end

  defp connect(_heap, set, 0) do
    set
  end

  defp connect(heap, set, num) do
    {_, from, to} = Heap.root(heap)
    connect(Heap.pop(heap), DisjointSet.union(set, from, to), num - 1)
  end

  defp distances_heap(connections) do
    Enum.reduce(connections, {Heap.new(:<), Enum.drop(connections, 1)}, fn connection,
                                                                           {acc, following_boxes} ->
      {Enum.map(following_boxes, fn x ->
         # {distance, from, to}
         {euclidean_distance(connection, x), connection, x}
       end)
       |> Enum.reduce(acc, fn x, acc_heap -> Heap.push(acc_heap, x) end),
       Enum.drop(following_boxes, 1)}
    end)
    |> elem(0)
  end

  defp euclidean_distance({x1, y1, z1}, {x2, y2, z2}) do
    :math.sqrt(:math.pow(x1 - x2, 2) + :math.pow(y1 - y2, 2) + :math.pow(z1 - z2, 2))
  end

  defp parse(input) do
    String.trim(input)
    |> String.split("\n")
    |> Enum.map(fn line ->
      line
      |> String.split(",")
      |> Enum.map(&parse_int!/1)
      |> List.to_tuple()
    end)
  end

  defp parse_int!(num) do
    {x, _} = Integer.parse(num)
    x
  end
end
