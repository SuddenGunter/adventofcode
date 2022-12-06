defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "Task1=1" do
    assert Task1.solution("mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 7
  end

  test "Task1-2" do
    assert Task1.solution("bvwbjplbgvbhsrlpgdmjqwftvncz") == 5
  end

  test "Task1-3" do
    assert Task1.solution("nppdvjthqldpwncqszvftbrmjlhg") == 6
  end

  test "Task1-4" do
    assert Task1.solution("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 10
  end

  test "Task1-5" do
    assert Task1.solution("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 11
  end
end
