defmodule Task2 do
  alias Collections.DisjointSet
  @spec solution(String.t()) :: integer()
  def solution(input) do
    boxes = parse(input)

    distances_heap(boxes)
    |> connect(DisjointSet.new(), boxes)
  end

  defp connect(heap, set, all_nodes) do
    {_, from, to} = Heap.root(heap)
    new_heap = Heap.pop(heap)
    new_set = DisjointSet.union(set, from, to)

    {root, _} = Collections.DisjointSet.find(new_set, hd(all_nodes))

    all_connected? =
      Enum.all?(all_nodes, fn x ->
        {x_root, _} = Collections.DisjointSet.find(new_set, x)
        x_root == root
      end)

    case all_connected? do
      true ->
        {from_x, _, _} = from
        {to_x, _, _} = to
        from_x * to_x

      false ->
        connect(new_heap, new_set, all_nodes)
    end
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
