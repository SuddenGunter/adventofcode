defmodule Task1 do
  def solution(data) do
    {total_horizontal, total_vertical} = Enum.reduce(
      data,
      {0, 0},
      fn x, {h, v} ->
        case x  do
          {:forward, distance} -> {h + distance, v}
          {:up, distance} -> {h, v - distance}
          {:down, distance} -> {h, v + distance}
        end
      end)

    total_horizontal * total_vertical
  end
end
