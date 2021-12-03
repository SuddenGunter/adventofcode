use Bitwise

defmodule Task1 do
  def replicate(n, x), do: for(i <- 0..n, i > 0, do: x)

  def solution(inputs) do
    oneAtEachIndex =
      Enum.reduce(
        inputs,
        Sh.replicate(String.length(Enum.at(inputs, 1)), 0),
        fn x, res ->
          Stream.with_index(res, 0)
          |> Enum.map(fn {el, index} ->
            el +
              if String.at(x, index) == "1" do
                1
              else
                0
              end
          end)
        end
      )

    gamma =
      oneAtEachIndex
      |> Enum.map(fn x ->
        if x > Enum.count(inputs) / 2 do
          1
        else
          0
        end
      end)
      |> Enum.reduce(fn x, acc -> (acc <<< 1) + x end)

    epsilon =
      oneAtEachIndex
      |> Enum.map(fn x ->
        if x > Enum.count(inputs) / 2 do
          0
        else
          1
        end
      end)
      |> Enum.reduce(fn x, acc -> (acc <<< 1) + x end)

    gamma * epsilon
  end
end
