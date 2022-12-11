defmodule Task1 do
  @spec solution([Monkey.t()]) :: integer()
  def solution(monkeys) do
    Solver.solution(monkeys, 20, fn x -> div(x, 3) end)
  end
end
