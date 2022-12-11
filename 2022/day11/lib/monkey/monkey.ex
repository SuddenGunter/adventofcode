defmodule Monkey do
  @enforce_keys [:id, :items, :operation, :test, :test_number]
  defstruct id: 0, items: [], operation: nil, test: nil, test_number: 0

  @type t :: %__MODULE__{
          id: integer(),
          items: [integer()],
          operation: (integer() -> integer()),
          test: (integer() -> integer())
        }

  @spec new(
          integer(),
          [integer()],
          (integer() -> integer()),
          (integer() -> integer()),
          integer()
        ) :: Monkey.t()
  def new(id, items, operation, test, test_number) do
    %Monkey{
      id: id,
      items: items,
      operation: operation,
      test: test,
      test_number: test_number
    }
  end
end
