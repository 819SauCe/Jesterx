import { useEffect, useState, FormEvent } from "react";
import { useNavigate, useParams } from "react-router-dom";
import styles from "../styles/pages/PageEdit.module.scss";
import { get, put } from "../utils/api";

type PageMeta = {
    id: string;
    tenant_id: string;
    name: string;
    page_id: string;
    domain?: string;
    theme_id?: string;
    created_at: string;
    updated_at: string;
};

type RawPage = {
    id: string;
    tenant_id: string;
    page_id: string;
    svelte: string;
    created_at: string;
    updated_at: string;
};

export function PageEdit() {
    const { pageId } = useParams<{ pageId: string }>();
    const navigate = useNavigate();

    const [loading, setLoading] = useState(true);
    const [saving, setSaving] = useState(false);
    const [error, setError] = useState("");
    const [success, setSuccess] = useState("");

    const [pageMeta, setPageMeta] = useState<PageMeta | null>(null);
    const [name, setName] = useState("");
    const [domain, setDomain] = useState("");
    const [svelteCode, setSvelteCode] = useState("");

    useEffect(() => {
        if (!pageId) {
            setError("Página inválida.");
            setLoading(false);
            return;
        }

        (async () => {
            try {
                const [metaRes, rawRes] = await Promise.all([
                    get<PageMeta>(`/v1/pages/${pageId}`),
                    get<RawPage>(`/v1/pages/${pageId}/raw`),
                ]);

                if (metaRes.data) {
                    setPageMeta(metaRes.data);
                    setName(metaRes.data.name);
                    setDomain(metaRes.data.domain || "");
                }

                if (rawRes.data) {
                    setSvelteCode(rawRes.data.svelte);
                }
            } catch (err: any) {
                const status = err?.status ?? err?.response?.status;
                if (status === 404) {
                    setError("Página não encontrada.");
                    return;
                }
                setError(err?.message || "Erro ao carregar página.");
            } finally {
                setLoading(false);
            }
        })();
    }, [pageId]);

    async function handleSubmit(e: FormEvent) {
        e.preventDefault();

        if (!pageId) return;

        setError("");
        setSuccess("");
        setSaving(true);

        try {
            await put(`/v1/pages/${pageId}`, {
                name,
                domain: domain || undefined,
                template: svelteCode,
            });

            setSuccess("Página atualizada com sucesso.");
        } catch (err: any) {
            setError(err?.message || "Erro ao salvar página.");
        } finally {
            setSaving(false);
        }
    }

    if (loading) {
        return (
            <main className={styles.main}>
                <div className={styles.loading}>Carregando editor...</div>
            </main>
        );
    }

    if (error && !pageMeta) {
        return (
            <main className={styles.main}>
                <p className={styles.error}>{error}</p>
                <button type="button" className={styles.secondaryButton} onClick={() => navigate("/pages")}>
                    Voltar para minhas páginas
                </button>
            </main>
        );
    }

    return (
        <main className={styles.main}>
            <header className={styles.header}>
                <h1 className={styles.title}>Editar página</h1>
                {pageMeta && (
                    <p className={styles.meta}>
                        ID: {pageMeta.id} • Slug: {pageMeta.page_id}
                    </p>
                )}
            </header>

            <form className={styles.form} onSubmit={handleSubmit}>
                {error && <div className={styles.errorBox}>{error}</div>}
                {success && <div className={styles.successBox}>{success}</div>}

                <label className={styles.field}>
                    <span className={styles.label}>Nome da página</span>
                    <input
                        className={styles.input}
                        type="text"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        disabled={saving}
                        required
                    />
                </label>

                <label className={styles.field}>
                    <span className={styles.label}>Domínio (opcional)</span>
                    <input
                        className={styles.input}
                        type="text"
                        value={domain}
                        onChange={(e) => setDomain(e.target.value)}
                        disabled={saving}
                        placeholder="ex: meusite.com.br"
                    />
                </label>

                <label className={styles.field}>
                    <span className={styles.label}>Código Svelte</span>
                    <textarea
                        className={styles.textarea}
                        value={svelteCode}
                        onChange={(e) => setSvelteCode(e.target.value)}
                        disabled={saving}
                        rows={18}
                    />
                </label>

                <div className={styles.actions}>
                    <button
                        type="button"
                        className={styles.secondaryButton}
                        onClick={() => navigate("/pages")}
                        disabled={saving}
                    >
                        Cancelar
                    </button>
                    <button type="submit" className={styles.primaryButton} disabled={saving}>
                        {saving ? "Salvando..." : "Salvar alterações"}
                    </button>
                </div>
            </form>
        </main>
    );
}
