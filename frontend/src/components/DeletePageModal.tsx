import { FormEvent, useState } from "react";
import styles from "../styles/pages/Pages.module.scss";
import { del } from "../utils/api";

type DeletePageModalProps = {
    open: boolean;
    pageName: string;
    pageId: string;
    onClose: () => void;
    onDeleted: () => Promise<void> | void;
};

export function DeletePageModal({ open, pageName, pageId, onClose, onDeleted }: DeletePageModalProps) {
    const [deleteConfirm, setDeleteConfirm] = useState("");
    const [deleting, setDeleting] = useState(false);
    const [deleteError, setDeleteError] = useState("");

    if (!open) return null;

    async function handleConfirmDelete(e: FormEvent) {
        e.preventDefault();

        const normalized = deleteConfirm.trim().toLowerCase();
        if (normalized !== "delete my page") {
            setDeleteError('Você precisa digitar exatamente: "delete my page".');
            return;
        }

        setDeleting(true);
        setDeleteError("");

        try {
            await del(`/v1/pages/${pageId}`);
            setDeleteConfirm("");
            await onDeleted();
            onClose();
        } catch (err: any) {
            const status = err?.status ?? err?.response?.status;
            if (status === 404) {
                setDeleteError("Página não encontrada.");
                return;
            }
            setDeleteError(err?.message || "Erro ao excluir página");
        } finally {
            setDeleting(false);
        }
    }

    return (
        <div className={styles.modalOverlay} onClick={onClose}>
            <div className={styles.modalContent} onClick={(e) => e.stopPropagation()}>
                <div className={styles.modalHeader}>
                    <h2>Excluir página</h2>
                    <p>Tem certeza que deseja excluir a página "{pageName}"? Esta ação não pode ser desfeita.</p>
                </div>

                <form onSubmit={handleConfirmDelete} className={styles.siteForm}>
                    {deleteError && <div className={styles.errorBox}>{deleteError}</div>}

                    <label>
                        Para confirmar, digite: <strong>delete my page</strong>
                        <input
                            type="text"
                            value={deleteConfirm}
                            onChange={(e) => setDeleteConfirm(e.target.value)}
                            required
                            disabled={deleting}
                            placeholder="delete my page"
                        />
                    </label>

                    <div className={styles.modalActions}>
                        <button type="button" className={styles.secondaryBtn} onClick={onClose} disabled={deleting}>
                            Cancelar
                        </button>
                        <button type="submit" className={styles.dangerBtn} disabled={deleting}>
                            {deleting ? "Excluindo..." : "Excluir página"}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
}
