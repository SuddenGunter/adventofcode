defmodule Day8Test do
  use ExUnit.Case

  test "Task 1 solution test" do
    assert Task1.solution([
             "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"
           ]) == 2
  end

  test "Task 2 solution test" do
    assert Task2.solution([
             "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"
           ]) == 8394
  end
end
