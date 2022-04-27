defmodule Task2Test do
  use ExUnit.Case
  doctest Task2

  test "duplicates over letter are nice" do
    assert Task2.solution(["xxyxx", "qjhvhtzxzqqjkmpb"]) == 2
  end

  test "uurcxstgmygtbstg is not nice" do
    assert Task2.solution(["uurcxstgmygtbstg", "ieodomkazucvgmuy"]) == 0
  end
end
