defmodule Task2Test do
  use ExUnit.Case
  doctest Task2

  test "Task2=1" do
    assert Task2.solution("mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 19
  end

  test "Task2-2" do
    assert Task2.solution("bvwbjplbgvbhsrlpgdmjqwftvncz") == 23
  end

  test "Task2-3" do
    assert Task2.solution("nppdvjthqldpwncqszvftbrmjlhg") == 23
  end

  test "Task2-4" do
    assert Task2.solution("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 29
  end

  test "Task2-5" do
    assert Task2.solution("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 26
  end
end
