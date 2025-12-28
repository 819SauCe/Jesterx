import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import styles from "../styles/pages/PagePreview.module.scss";
import { get } from "../utils/api";

type RawPage = {
    id: string;
    tenant_id: string;
    page_id: string;
    svelte: string;
    created_at: string;
    updated_at: string;
};

export function PagePreview() {
    const { pageId } = useParams<{ pageId: string }>();
    const navigate = useNavigate();

    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");
    const [page, setPage] = useState<RawPage | null>(null);

    useEffect(() => {
        if (!pageId) {
            setError("Página inválida.");
            setLoading(false);
            return;
        }

        (async () => {
            try {
                const res = await get<RawPage>(`/v1/pages/${pageId}/raw`);
                if (res.data) {
                    setPage(res.data);
                } else {
                    setError("Página não encontrada.");
                }
            } catch (err: any) {
                const status = err?.status ?? err?.response?.status;
                if (status === 404) {
                    setError("Página não encontrada.");
                    return;
                }
                if (status === 401) {
                    navigate("/login");
                    return;
                }
                setError(err?.message || "Erro ao carregar página.");
            } finally {
                setLoading(false);
            }
        })();
    }, [pageId, navigate]);

    if (loading) {
        return (
            <main className={styles.main}>
                <div className={styles.center}>
                    <p>Carregando página...</p>
                </div>
            </main>
        );
    }

    if (error || !page) {
        return (
            <main className={styles.main}>
                <div className={styles.center}>
                    <p className={styles.error}>{error || "Página não encontrada."}</p>
                    <button type="button" className={styles.backButton} onClick={() => navigate("/pages")}>
                        Voltar para minhas páginas
                    </button>
                </div>
            </main>
        );
    }

    return (
        <main className={styles.fullPage}>
            <iframe
                title={page.page_id}
                className={styles.iframe}
                srcDoc={page.svelte}
                sandbox="allow-same-origin allow-scripts allow-forms allow-popups"
            />
        </main>
    );
}
