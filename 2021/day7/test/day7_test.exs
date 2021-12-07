defmodule Day7Test do
  use ExUnit.Case

  test "Task 1 solution test" do
    assert Task1.solution([16, 1, 2, 0, 4, 2, 7, 1, 2, 14]) == 37
  end

  test "Task 2 solution test" do
    assert Task2.solution([16, 1, 2, 0, 4, 2, 7, 1, 2, 14]) == 168
  end
end
