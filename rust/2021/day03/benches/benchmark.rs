use criterion::{black_box, criterion_group, criterion_main, Criterion};
use day03::{part1, part2};

criterion_group!(benches, benchmark_part1, benchmark_part2);
criterion_main!(benches);

fn benchmark_part1(c: &mut Criterion) {
    c.bench_function("part-1", |b| {
        b.iter(|| black_box(part1(black_box(include_str!("../input.txt").to_string()))));
    });
}

fn benchmark_part2(c: &mut Criterion) {
    c.bench_function("part-2", |b| {
        b.iter(|| black_box(part2(black_box(include_str!("../input.txt").to_string()))));
    });
}
