defmodule Day7Test do
  use ExUnit.Case
  doctest Task1

  test "Task 1 solution test" do
    assert Task1.solution([16,1,2,0,4,2,7,1,2,14]) == 37
  end
end
