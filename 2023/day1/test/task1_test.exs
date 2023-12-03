defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "Task1-1" do
    assert Task1.solution(["1abc2"]) == 12
  end

  test "Task1-2" do
    assert Task1.solution(["a1b2c3d4e5f"]) == 15
  end

  test "Task1-3" do
    assert Task1.solution(["treb7uchet"]) == 77
  end

  test "Task1-4" do
    assert Task1.solution([
             "1abc2",
             "pqr3stu8vwx",
             "a1b2c3d4e5f",
             "treb7uchet"
           ]) == 142
  end
end
