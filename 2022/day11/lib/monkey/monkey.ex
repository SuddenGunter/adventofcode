defmodule Monkey do
  @enforce_keys [:id, :items, :operation, :test]
  defstruct id: 0, items: [], operation: nil, test: nil

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
          (integer() -> integer())
        ) :: Monkey.t()
  def new(id, items, operation, test) do
    %Monkey{
      id: id,
      items: items,
      operation: operation,
      test: test
    }
  end
end
