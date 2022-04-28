defmodule Task1Test do
  use ExUnit.Case
  doctest Task1

  test "value is evaled to value" do
    assert Task1.solution(%{"a" => %Signal{left: 42, gate: :value}}) == 42
  end

  test "embedded value is evaled to value" do
    assert Task1.solution(%{
             "a" => %Signal{left: "b", gate: :value},
             "b" => %Signal{left: 42, gate: :value}
           }) == 42
  end

  test "and" do
    assert Task1.solution(%{
             "a" => %Signal{left: "tmp1", right: "tmp2", gate: :and},
             "tmp1" => %Signal{left: 1, gate: :value},
             "tmp2" => %Signal{left: 3, gate: :value}
           }) == 1
  end

  test "or" do
    assert Task1.solution(%{
             "a" => %Signal{left: "tmp1", right: "tmp2", gate: :or},
             "tmp1" => %Signal{left: 1, gate: :value},
             "tmp2" => %Signal{left: 3, gate: :value}
           }) == 3
  end

  test "lshift" do
    assert Task1.solution(%{
             "a" => %Signal{left: "tmp1", right: "tmp2", gate: :lshift},
             "tmp1" => %Signal{left: 1, gate: :value},
             "tmp2" => %Signal{left: 1, gate: :value}
           }) == 2
  end

  test "rshift" do
    assert Task1.solution(%{
             "a" => %Signal{left: "tmp1", right: "tmp2", gate: :rshift},
             "tmp1" => %Signal{left: 1, gate: :value},
             "tmp2" => %Signal{left: 1, gate: :value}
           }) == 0
  end
end
