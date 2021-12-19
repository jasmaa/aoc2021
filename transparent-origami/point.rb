class Point
  attr_accessor :x
  attr_accessor :y

  def initialize(x, y)
    @x = x
    @y = y
  end

  def hash
    "#{x},#{y}".hash
  end

  def eql? other
    self.hash == other.hash
  end
end