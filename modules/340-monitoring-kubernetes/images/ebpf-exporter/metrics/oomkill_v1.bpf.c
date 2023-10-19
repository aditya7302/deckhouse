#include <uapi/linux/ptrace.h>
#include <linux/oom.h>
#include <linux/memcontrol.h>
// we'll use "BPF_PERF_OUTPUT" map type here to avoid unbound cardinality
BPF_PERF_OUTPUT(events);
struct data_t {
    u64 cgroup_id;
    u8 global_oom;
};
void count_ooms(struct pt_regs *ctx, struct oom_control *oc, const char *message) {
    struct data_t data = {};
    struct mem_cgroup *mcg = oc->memcg;
    if (!mcg) {
        data.global_oom = 1;
        events.perf_submit(ctx, &data, sizeof(data));
        return;
    }
    data.cgroup_id = mcg->css.cgroup->kn->id.id;
    events.perf_submit(ctx, &data, sizeof(data));
}

