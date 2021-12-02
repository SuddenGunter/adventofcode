defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "Task1: dive (test case)" do
    assert Task1.solution([
             {:forward, 5},
             {:down, 5},
             {:forward, 8},
             {:up, 3},
             {:down, 8},
             {:forward, 2}
           ]) == 150
  end
end
