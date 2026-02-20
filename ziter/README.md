# ziter

Iterator utilities for Go's `iter.Seq[V]` and `iter.Seq2[K, V]` types.

## Functions

### Create

| Function | Signature | Description |
|----------|-----------|-------------|
| `Single` | `Single[V any](v V) iter.Seq[V]` | Sequence with exactly one element |
| `Single2` | `Single2[K, V any](k K, v V) iter.Seq2[K, V]` | Sequence with exactly one key-value pair |

### Transform

| Function | Signature | Description |
|----------|-----------|-------------|
| `Map` | `Map[V1, V2 any](seq iter.Seq[V1], f func(V1) V2) iter.Seq[V2]` | Transform each element |
| `Map2` | `Map2[K1, V1, K2, V2 any](seq iter.Seq2[K1, V1], f func(K1, V1) (K2, V2)) iter.Seq2[K2, V2]` | Transform each key-value pair |
| `MapKey` | `MapKey[K1, V, K2 any](seq iter.Seq2[K1, V], f func(K1) K2) iter.Seq2[K2, V]` | Transform keys |
| `MapKey2` | `MapKey2[K1, V, K2 any](seq iter.Seq2[K1, V], f func(K1, V) K2) iter.Seq2[K2, V]` | Transform keys (with access to value) |
| `MapValue` | `MapValue[K, V1, V2 any](seq iter.Seq2[K, V1], f func(V1) V2) iter.Seq2[K, V2]` | Transform values |
| `MapValue2` | `MapValue2[K, V1, V2 any](seq iter.Seq2[K, V1], f func(K, V1) V2) iter.Seq2[K, V2]` | Transform values (with access to key) |

### Flatten

| Function | Signature | Description |
|----------|-----------|-------------|
| `Flatten` | `Flatten[V1, V2 any](seq iter.Seq[V1], f func(V1) []V2) iter.Seq[V2]` | Map each element to a slice, yield all results |
| `FlattenKeys` | `FlattenKeys[K1, V, K2 any](seq iter.Seq2[K1, V], f func(K1) []K2) iter.Seq2[K2, V]` | Expand keys, pairing each with the original value |
| `FlattenValues` | `FlattenValues[K, V1, V2 any](seq iter.Seq2[K, V1], f func(V1) []V2) iter.Seq2[K, V2]` | Expand values, pairing each with the original key |

### Filter

| Function | Signature | Description |
|----------|-----------|-------------|
| `Filter` | `Filter[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V]` | Keep elements matching predicate |
| `Filter2` | `Filter2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V]` | Keep pairs matching predicate |
| `FilterKey` | `FilterKey[K, V any](seq iter.Seq2[K, V], f func(K) bool) iter.Seq2[K, V]` | Keep pairs where key matches |
| `FilterValue` | `FilterValue[K, V any](seq iter.Seq2[K, V], f func(V) bool) iter.Seq2[K, V]` | Keep pairs where value matches |
| `Dedup` | `Dedup[V comparable](seq iter.Seq[V]) iter.Seq[V]` | Remove duplicate elements |
| `Dedup2` | `Dedup2[K, V comparable](seq iter.Seq2[K, V]) iter.Seq2[K, V]` | Remove duplicate pairs |

### Slice / Take / Drop

| Function | Signature | Description |
|----------|-----------|-------------|
| `Take` | `Take[V any](seq iter.Seq[V], n int) iter.Seq[V]` | First n elements |
| `Take2` | `Take2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V]` | First n pairs |
| `Drop` | `Drop[V any](seq iter.Seq[V], n int) iter.Seq[V]` | Skip first n elements |
| `Drop2` | `Drop2[K, V any](seq iter.Seq2[K, V], n int) iter.Seq2[K, V]` | Skip first n pairs |
| `TakeWhile` | `TakeWhile[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V]` | Take while predicate holds |
| `TakeWhile2` | `TakeWhile2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V]` | Take pairs while predicate holds |
| `DropWhile` | `DropWhile[V any](seq iter.Seq[V], f func(V) bool) iter.Seq[V]` | Skip while predicate holds, then yield rest |
| `DropWhile2` | `DropWhile2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) iter.Seq2[K, V]` | Skip pairs while predicate holds |
| `Chunk` | `Chunk[V any](seq iter.Seq[V], n int) iter.Seq[[]V]` | Group into slices of up to n elements |

### Combine

| Function | Signature | Description |
|----------|-----------|-------------|
| `Concat` | `Concat[V any](seqs ...iter.Seq[V]) iter.Seq[V]` | Concatenate sequences |
| `Concat2` | `Concat2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V]` | Concatenate pair sequences |
| `Zip` | `Zip[A, B any](seqA iter.Seq[A], seqB iter.Seq[B]) iter.Seq2[A, B]` | Combine two sequences element-wise into pairs |

### Split

| Function | Signature | Description |
|----------|-----------|-------------|
| `Split` | `Split[V any](seq iter.Seq[V], f func(V) bool) (iter.Seq[V], iter.Seq[V])` | Partition into matching and non-matching |
| `Split2` | `Split2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) (iter.Seq2[K, V], iter.Seq2[K, V])` | Partition pairs by predicate |
| `SplitKey` | `SplitKey[K, V any](seq iter.Seq2[K, V], f func(K) bool) (iter.Seq2[K, V], iter.Seq2[K, V])` | Partition pairs by key |
| `SplitValue` | `SplitValue[K, V any](seq iter.Seq2[K, V], f func(V) bool) (iter.Seq2[K, V], iter.Seq2[K, V])` | Partition pairs by value |

### Convert between Seq and Seq2

| Function | Signature | Description |
|----------|-----------|-------------|
| `ToSeq2` | `ToSeq2[E, K, V any](seq iter.Seq[E], f func(E) (K, V)) iter.Seq2[K, V]` | Convert Seq to Seq2 via function |
| `ToSeq1` | `ToSeq1[K, V, E any](seq iter.Seq2[K, V], f func(K, V) E) iter.Seq[E]` | Convert Seq2 to Seq via function |
| `KeyBy` | `KeyBy[V, K2 any](seq iter.Seq[V], f func(V) K2) iter.Seq2[K2, V]` | Seq to Seq2, deriving key from element |
| `ValueBy` | `ValueBy[V, V2 any](seq iter.Seq[V], f func(V) V2) iter.Seq2[V, V2]` | Seq to Seq2, deriving value from element |
| `Enumerate` | `Enumerate[V any](seq iter.Seq[V]) iter.Seq2[int, V]` | Seq to Seq2 with zero-based index as key |
| `Keys` | `Keys[K, V any](seq iter.Seq2[K, V]) iter.Seq[K]` | Extract keys |
| `Values` | `Values[K, V any](seq iter.Seq2[K, V]) iter.Seq[V]` | Extract values |

### Reduce (terminal)

| Function | Signature | Description |
|----------|-----------|-------------|
| `Reduce` | `Reduce[V any](seq iter.Seq[V], f func(V, V) V) (V, bool)` | Accumulate using first element as initial value |
| `Reduce2` | `Reduce2[K, V any](seq iter.Seq2[K, V], f func(K, V, K, V) (K, V)) (K, V, bool)` | Accumulate pairs using first pair as initial value |
| `Aggregate` | `Aggregate[V, A any](seq iter.Seq[V], init A, f func(A, V) A) A` | Fold into a different type |
| `Aggregate2` | `Aggregate2[K, V, A any](seq iter.Seq2[K, V], init A, f func(A, K, V) A) A` | Fold pairs into a different type |
| `Count` | `Count[V any](seq iter.Seq[V]) int` | Number of elements |
| `Count2` | `Count2[K, V any](seq iter.Seq2[K, V]) int` | Number of pairs |
| `Max` | `Max[V cmp.Ordered](seq iter.Seq[V]) (V, bool)` | Maximum of ordered elements |
| `Min` | `Min[V cmp.Ordered](seq iter.Seq[V]) (V, bool)` | Minimum of ordered elements |
| `MaxFunc` | `MaxFunc[V any](seq iter.Seq[V], cmp func(V, V) int) (V, bool)` | Maximum using custom comparator |
| `MinFunc` | `MinFunc[V any](seq iter.Seq[V], cmp func(V, V) int) (V, bool)` | Minimum using custom comparator |

### Find (terminal)

| Function | Signature | Description |
|----------|-----------|-------------|
| `FindAny` | `FindAny[V any](seq iter.Seq[V]) (V, bool)` | Any element, or false if empty |
| `FindAny2` | `FindAny2[K, V any](seq iter.Seq2[K, V]) (K, V, bool)` | Any pair, or false if empty |
| `FindFirst` | `FindFirst[V any](seq iter.Seq[V], f func(V) bool) (V, bool)` | First element matching predicate |
| `FindFirst2` | `FindFirst2[K, V any](seq iter.Seq2[K, V], f func(K, V) bool) (K, V, bool)` | First pair matching predicate |
| `Exists` | `Exists[V any](seq iter.Seq[V]) bool` | Whether sequence is non-empty |
