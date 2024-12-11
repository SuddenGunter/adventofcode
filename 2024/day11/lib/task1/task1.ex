defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
   input
   |> String.split("\n", trim: true)
   |> Enum.at(0)
   |> String.split(" ", trim: true)
   |> Enum.map(&String.to_integer/1)
   |> IO.inspect( )
  end
end
