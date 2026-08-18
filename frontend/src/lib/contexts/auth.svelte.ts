export function createAuth() {
	let user = $state<string | null>(null);
	let loading = $state(false);

    return {
        get user() {
            return user;
        },

        get loading() {
            return loading;
        },
    }
}
