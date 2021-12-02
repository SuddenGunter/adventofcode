defmodule Task2 do
  def solution(data) do
    {total_horizontal, total_vertical, _} =
      Enum.reduce(
        data,
        {0, 0, 0},
        fn x, {h, v, aim} ->
          case x do
            {:forward, distance} -> {h + distance, v + aim * distance, aim}
            {:up, distance} -> {h, v, aim - distance}
            {:down, distance} -> {h, v, aim + distance}
          end
        end
      )

    total_horizontal * total_vertical
  end
end
