require 'set'
require './point'
require './instruction'

class Scenario
  attr_accessor :points
  attr_accessor :instructions

  def initialize(points, instructions)
    @points = points
    @instructions = instructions
    @instruction_idx = 0
  end

  def self.parse(payload)
    lines = payload.split(/\r?\n/)
    points = []
    instructions = []
    parsing_points = true
    lines.each do |line|
      if parsing_points
        if line.length == 0
          parsing_points = false
        else
          captures = line.match(/^(\d+),(\d+)$/).captures
          points << Point.new(captures[0].to_i, captures[1].to_i)
        end
      else
        captures = line.match(/^fold along (x|y)=(\d+)$/).captures
        instructions << Instruction.new(captures[0], captures[1].to_i)
      end
    end
    Scenario.new points, instructions
  end

  def steps_left?
    return @instruction_idx < @instructions.length
  end

  def step
    # Fold
    ins = @instructions[@instruction_idx]
    case ins.direction
    when 'x'
      @points.each do |p|
        if p.x > ins.value
          p.x = -(p.x - ins.value) + ins.value
        end
      end
    when 'y'
      @points.each do |p|
        if p.y > ins.value
          p.y = -(p.y - ins.value) + ins.value
        end
      end
    else
      raise 'invalid direction'
    end
    # Get rid of dupes
    @points = @points.to_set.to_a
    @instruction_idx += 1
  end

  def to_ascii
    origin_offset = self.class.find_min_offset @points
    offset_points = @points.map{ |p| Point.new p.x - origin_offset.x, p.y - origin_offset.y }
    max_bounds = self.class.find_max_offset offset_points
    grid = Array.new(max_bounds.y + 1, ".").map{ |_| Array.new(max_bounds.x + 1, ".") }
    offset_points.each do |p|
      grid[p.y][p.x] = "#"
    end
    grid.map{ |row| row.join("") }.join("\n")
  end

  private

  def self.find_min_offset(points)
    min_x = points.map{ |p| p.x }.min
    min_y = points.map{ |p| p.y }.min
    Point.new min_x, min_y
  end

  def self.find_max_offset(points)
    max_x = points.map{ |p| p.x }.max
    max_y = points.map{ |p| p.y }.max
    Point.new max_x, max_y
  end
end