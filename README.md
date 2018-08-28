# Go Kata Practice
  [![Build Status][travis-image]][travis-url]
  [![Build status](https://ci.appveyor.com/api/projects/status/94djahf3vnm1tk51?svg=true)](https://ci.appveyor.com/project/liuwill/go-kata-practice)
  [![Codecov branch][codecov-image]][codecov-url]
  [![Maintainability][codeclimate-image]][codeclimate-url]

JUST make some practice


## Docs

> [Say All Puzzle](./docs)

### Questions

-   [Sudoku](./docs/sudoku.md)
-   [Trapping Rain Water](./docs/trap_rain_water.md)
-   [Jump Game](./docs/jump_game.md)
-   [Jump Game II](./docs/jump_game_II.md)
-   [Multiply Strings](./docs/multiply_strings.md)
-   [Add Strings](./docs/add_strings.md)
-   [Swap Node In Pair](./docs/swap_node_in_pairs.md)
-   [Best Time to Buy and Sell Stock](./docs/best_time_to_buy_and_sell_stock.md)
-   [Best Time to Buy and Sell Stock II](./docs/best_time_to_buy_and_sell_stock_II.md)
-   [Remove Duplicates from Sorted Array](./docs/remove_duplicates_from_sorted_array.md)
-   [Numbers Of Island](./docs/numbers_of_island.md)
-   [House Robber](./docs/house_robber.md)
-   [Keys And Rooms](./docs/keys_and_rooms.md)
-   [Dungeon Game](./docs/dungeon_game.md)
-   [Count And Say](./docs/count_and_say.md)
-   [Gas Station](./docs/gas_station.md)
-   [Basic Calculator II](./docs/basic_calculator_II.md)
-   [Rotate Image](./docs/rotate_image.md)
-   [Sort Colors](./docs/sort_colors.md)
-   [Binary Gap](./docs/binary_gap.md)
-   [Advantage Shuffle](./docs/advantage_shuffle.md)
-   [Super Egg Drop](./docs/super_egg_drop.md)
-   [Uncommon Words from Two Sentences](./docs/uncommon_words_from_two_sentences.md)
-   [Boats to Save People](./docs/boats_to_save_people.md)
-   [Koko Eating Bananas](./docs/koko_eating_bananas.md)
-   [Largest Triangle Area](./docs/largest_triangle_area.md)
-   [Fair Candy Swap](./docs/fair_candy_swap.md)
-   [Find and Replace Pattern](./docs/find_and_replace_pattern.md)
-   [Minimum Number of Refueling Stops](./docs/minimum_number_of_refueling_stops.md)
-   [Masking Personal Information](./docs/masking_personal_information.md)

#### Design
-   [Design Twitter](./docs/design_twitter.md)


### Script
```shell
# 运行单元测试
make test

# 单元测试并且生成测试覆盖报告
make coverage

# 单元测试并且测试覆盖HTML页面
make coverhtml
```

#### 运行单个测试

run.test.sh脚本，参数一设置目录

```shell
./run.test.sh rain
```

#### docker 环境运行

```shell
# 创建docker容器
./docker.sh install

# 进入docker容器
./docker.sh enter

# 在容器中运行命令
cd /go/src/go-kata-practice
make coverhtml

# 或者
./docker.sh test
```

## License

  [MIT](./LICENSE)

[travis-image]: https://img.shields.io/travis/liuwill/go-kata-practice/master.svg?style=flat-square
[travis-url]: https://travis-ci.org/liuwill/go-kata-practice
[codecov-image]: https://img.shields.io/codecov/c/github/liuwill/go-kata-practice.svg?style=flat-square
[codecov-url]: https://codecov.io/gh/liuwill/go-kata-practice
[codeclimate-image]: https://api.codeclimate.com/v1/badges/356d7f0824e1b1e5d9ff/maintainability
[codeclimate-url]: https://codeclimate.com/github/liuwill/go-kata-practice/maintainability
