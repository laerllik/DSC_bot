import { type Context } from 'grammy';

export type CommandHandler = {
    cmd: string;
    handler: (ctx: Context) => void;
    description: string;
};
