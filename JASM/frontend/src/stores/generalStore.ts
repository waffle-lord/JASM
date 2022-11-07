import { writable } from 'svelte/store';
import { RefreshRate } from '../data/refreshRate';

export const refreshRate = writable(RefreshRate.NONE);